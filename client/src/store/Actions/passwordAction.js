import {
	ROOT_API,
	EDIT_PASSWORD
} from "./types"
import {
	message
} from "antd";

export function handleEdit(payload) {
	return (dispatch) => {
		dispatch({
			type: EDIT_PASSWORD,
			payload: payload
		})
	}
}

export function updateNewPassword(savePassword, pusher) {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/user/update/${employeeNumber}`, {
				method: 'PUT',
				body: JSON.stringify(savePassword)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					message.success(body)
					pusher('/profile')
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @updateNewPassword: ", error)
			})
	}
}