import {
	CREATE_LEAVE_REQUEST,
	CLEAR_FIELD
} from "../Actions/types"

const leaveState = {
	type_leave_id: null,
	reason: '',
	date_from: '',
	date_to: '',
	half_dates: [],
	back_on: null,
	contact_address: '',
	contact_number: ''
}

export default function leaveRequestReducer(state = leaveState, action) {
	switch (action.type) {
		case CREATE_LEAVE_REQUEST:
			return {
				...action.payload,
			}
		case CLEAR_FIELD:
			return {
				...leaveState
			}
		default:
			return state
	}
}