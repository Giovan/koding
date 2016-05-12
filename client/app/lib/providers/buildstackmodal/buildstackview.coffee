kd = require 'kd'
async = require 'async'
remote = require('app/remote').getInstance()
JView = require 'app/jview'
CredentialForm = require './credentialform'
showError = require 'app/util/showError'

module.exports = class BuildStackView extends kd.View

  JView.mixin @prototype

  constructor: (options = {}, data) ->

    options.cssClass = 'build-stack-view'
    super options, data

    @createCredentialView()
    @createRequirementsView()

    @buildButton = new kd.ButtonView
      title    : 'Build Stack'
      cssClass : 'GenericButton'
      loader   : yes
      callback : @bound 'onBuild'


  createCredentialView: ->

    { credentials } = @getData()

    @awsCredentialContainer = new kd.CustomScrollView
      cssClass : 'form-scroll-wrapper credential-wrapper'
    @awsCredentialForm = new CredentialForm {
      title : 'AWS Credential'
      selectionPlaceholder : 'Select credential...'
      selectionLabel : 'Credential Selection'
    }, credentials
    @awsCredentialContainer.wrapper.addSubView @awsCredentialForm


  createRequirementsView: ->

    { requirements } = @getData()

    @requirementsContainer = new kd.CustomScrollView
      cssClass : 'form-scroll-wrapper requirements-wrapper'

    return @setClass 'credential-only'  unless requirements.fields

    @requirementsForm = new CredentialForm {
      title : 'Requirements'
      selectionPlaceholder : 'Select from existing requirements...'
      selectionLabel : 'Requirement Selection'
    }, requirements
    @requirementsContainer.wrapper.addSubView @requirementsForm


  buildTitleAndDescription: ->

    { credentials, requirements } = @getData()

    if not credentials.items.length and not requirements.fields
      return {
        title       : 'Create Your First Credential'
        description : '''
          Your Credential provides Koding with all of the information it needs to build your Stack
        '''
      }

    return {
      title       : 'Select Credential and Fill the Requirements'
      description : '''
        Your stack requires AWS Credentials and a few requirements in order to boot
      '''
    }


  onBuild: ->

    validationQueue =
      credential    : helper.createFormValidationCallback @awsCredentialForm
      requirements  : helper.createFormValidationCallback @requirementsForm

    async.parallel validationQueue, (err, validationResults) =>
      return @buildButton.hideLoader()  if err

      resultQueue = [
        helper.createFormResultCallback validationResults.credential
        helper.createFormResultCallback validationResults.requirements
      ]

      async.series resultQueue, (err, identifiers) =>
        if err
          @buildButton.hideLoader()
          return showError err

        alert identifiers


  pistachio: ->

    { title, description } = @buildTitleAndDescription()

    """
      <div class="credentials-step">
        <header>
          <h1>Build Your Stack</h1>
        </header>
        <section class="main">
          <h2>#{title}</h2>
          <p>#{description}</p>
          {{> @awsCredentialContainer}}
          {{> @requirementsContainer}}
          <div class="clearfix"></div>
        </section>
        <footer>
          {{> @buildButton}}
        </footer>
      </div>
    """

  helper =

    createFormValidationCallback: (form) ->

      (next) ->

        return next()  unless form

        form.off  'FormValidationPassed'
        form.once 'FormValidationPassed', (result) ->
          next null, result

        form.off  'FormValidationFailed'
        form.once 'FormValidationFailed', -> next 'ValidationError'

        form.validate()


    createFormResultCallback: (validationResult) ->

      (next) ->

        return next()  unless validationResult

        { selectedItem, newData } = validationResult
        return next null, selectedItem  if selectedItem

        remote.api.JCredential.create newData, (err, credential) ->
          return next err  if err
          return next null, credential.identifier
