import {
	ROOT_API,
	REGISTER_USER,
	CLEAR_FIELD
} from "./types"
import {
	message
} from "antd";

export function formOnchange(payload) {
	return (dispatch) => {
		dispatch({
			type: REGISTER_USER,
			payload: payload
		})
	}
}

function clearField(msg) {
	return {
		type: CLEAR_FIELD
	}
}

export function registerUser(payload, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/register`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					message.success(body)
					dispatch(clearField)
					pusher('/admin')
				} else if (error === "Type request malform") {
					let errMsg = 'Error empty field!'
					message.error(errMsg)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @SumbitSignUp: ", error)
			})
	}
}