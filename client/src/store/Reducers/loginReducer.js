import {
	LOGIN_FORM_ONCHANGE,
	CLEAR_FIELD
} from "../Actions/types"

const loginState = {
	email: '',
	password: ''
}

export default function loginReducer(state = loginState, action) {
	switch (action.type) {
		case LOGIN_FORM_ONCHANGE:
			return {
				...action.payload
			}
		case CLEAR_FIELD:
			return {
				...loginState,
			}
		default:
			return state
	}
}