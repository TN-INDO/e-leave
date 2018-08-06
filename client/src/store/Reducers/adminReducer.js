import {
	FETCH_USER,
	DELETE_USER,
	FETCH_LEAVE_PENDING,
	FETCH_LEAVE_APPROVE,
	FETCH_LEAVE_REJECT,
	CANCEL_LEAVE_REQUEST,
	FETCH_LEAVE_BALANCES
} from "../Actions/types"

let adminState = {
	loading: true,
	users: [],
	leaves: [],
	balances: [],
}

export default function adminReducer(state = adminState, action) {
	switch (action.type) {
		case FETCH_USER:
			return {
				...action.payload
			}
		case DELETE_USER:
			return {
				...action.payload
			}
		case FETCH_LEAVE_PENDING:
			return {
				...action.payload
			}
		case FETCH_LEAVE_APPROVE:
			return {
				...action.payload
			}
		case FETCH_LEAVE_REJECT:
			return {
				...action.payload
			}
		case CANCEL_LEAVE_REQUEST:
			return {
				...action.payload
			}
		case FETCH_LEAVE_BALANCES:
			return {
				...action.payload
			}
		default:
			return state
	}
}