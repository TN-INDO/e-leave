import {
	REGISTER_USER,
	CLEAR_FIELD
} from "../Actions/types"


// function makeid() {
// 	var x = 10;
// 	var text = "";
// 	var possible =
// 		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*?";

// 	for (var i = 0; i < 7; i++)
// 		text += possible.charAt(Math.floor(Math.random() * possible.length));

// 	return text;
// 	// setTimeout(makeid, x * 1000);
// }

const registerState = {
	employee_number: '',
	name: '',
	gender: '',
	position: '',
	start_working_date: '',
	mobile_phone: '',
	email: '',
	password: '',
	role: '',
	supervisor_id: null
}

export default function registerReducer(state = registerState, action) {
	switch (action.type) {
		case REGISTER_USER:
			return {
				...action.payload
			}
		case CLEAR_FIELD:
			return {
				...registerState
			}
		default:
			return state
	}
}