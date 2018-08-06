import {
	GET_PROFILE
} from "../Actions/types"

const profileState = {
	loading: true,
	user: []
}

export default function profileReducer(state = profileState, action) {
	switch (action.type) {
		case GET_PROFILE:
			return {
				...action.payload
			}
		default:
			return state
	}
}