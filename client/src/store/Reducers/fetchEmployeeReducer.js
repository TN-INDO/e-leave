import {
	FETCH_REQUEST_PENDING,
	DELETE_REQUEST_PENDING,
	FETCH_REQUEST_APPROVE,
	FETCH_REQUEST_REJECT,
} from "../Actions/types"

const leaveRequest = {
	loading: true,
	leaves: []
}

export default function fetchEmployeeReducer(state = leaveRequest, action) {
	switch (action.type) {
		case FETCH_REQUEST_PENDING:
			return {
				...action.payload
			}
		case 'UPDATE_REQUEST_PENDING':
			return {
				...action.payload
			}
		case DELETE_REQUEST_PENDING:
			return {
				...action.payload
			}
		case FETCH_REQUEST_APPROVE:
			return {
				...action.payload
			}
		case FETCH_REQUEST_REJECT:
			return {
				...action.payload
			}
		default:
			return state
	}
}