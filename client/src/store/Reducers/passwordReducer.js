import {
	EDIT_PASSWORD
} from "../Actions/types"

const passwordState = {
	old_password: "",
	new_password: "",
	confirm_password: ""
}

export default function passwordReducer(state = passwordState, action) {
	switch (action.type) {
		case EDIT_PASSWORD:
			return {
				...action.payload
			}
		default:
			return state
	}
}