kd = require 'kd'
actionTypes = require './actiontypes'

dispatch = (args...) -> kd.singletons.reactor.dispatch args...

FormValuesStore = require './stores/HomeTeamBillingFormValuesStore'

module.exports =
  stores: [ FormValuesStore ]

  actions:
    setValue: (type, value) -> dispatch actionTypes.SET_TEAM_BILLING_INPUT_VALUE, { type, value }
    resetValues: -> dispatch actionTypes.RESET_TEAM_BILLING_INPUT_VALUES

  getters: require './getters'

  register: (reactor) -> reactor.registerStores @stores

