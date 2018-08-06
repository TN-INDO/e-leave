import {
	FETCH_EDIT_USER,
	EDIT_USER
} from "../Actions/types"

const userState = {
	loading: true,
	user: {
		employee_number: '',
		name: '',
		gender: '',
		position: '',
		start_working_date: '',
		mobile_phone: '',
		email: '',
		role: '',
		supervisor_id: ''
	}
}

export default function editUserReducer(state = userState, action) {
	switch (action.type) {
		case FETCH_EDIT_USER:
			return {
				...action.payload
			}
		case EDIT_USER:
			return {
				...action.payload
			}
		default:
			return state
	}
}