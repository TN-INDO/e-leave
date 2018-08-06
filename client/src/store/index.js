import {
	createStore,
	combineReducers,
	applyMiddleware
} from 'redux'
import thunk from 'redux-thunk';

import loginReducer from './Reducers/loginReducer'
import resetPasswordReducer from './Reducers/resetPasswordReducer'

import adminReducer from './Reducers/adminReducer'
import registerReducer from './Reducers/registerReducer'
import editUserReducer from './Reducers/editUserReducer'

import leaveRequestReducer from './Reducers/leaveRequestReducer'
import editRequestReducer from './Reducers/editRequestReducer'

import fetchDirectorReducer from './Reducers/fetchDirectorReducer'
import fetchSupervisorReducer from './Reducers/fetchSupervisorReducer'
import fetchEmployeeReducer from './Reducers/fetchEmployeeReducer'

import fetchUserSummaryReducer from './Reducers/fetchUserSummaryReducer'
import profileReducer from './Reducers/profileReducer'
import passwordReducer from './Reducers/passwordReducer'

import fetchTypeLeaveReducer from './Reducers/fetchTypeLeaveReducer'
import AddSupervisorReducer from './Reducers/AddSupervisorReducer'


const appStore = combineReducers({
	loginReducer,
	resetPasswordReducer,

	adminReducer,
	registerReducer,
	editUserReducer,

	leaveRequestReducer,
	editRequestReducer,

	fetchDirectorReducer,
	fetchSupervisorReducer,
	fetchEmployeeReducer,

	fetchUserSummaryReducer,
	profileReducer,
	passwordReducer,

	fetchTypeLeaveReducer,
	AddSupervisorReducer,
})

const store = createStore(
	appStore,
	window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
	applyMiddleware(thunk)
)
export default store