import {
	ROOT_API,
	FETCH_EDIT_USER,
	EDIT_USER
} from "./types"
import {
	message
} from "antd";

function fetchUser(payload) {
	return {
		type: FETCH_EDIT_USER,
		payload: payload
	}
}

function userEditing(payload) {
	return {
		type: EDIT_USER,
		payload: payload
	}
}

export function handleEdit(newUser) {
	return (dispatch) => {
		let payload = {
			loading: false,
			user: {
				...newUser
			}
		}
		dispatch(userEditing(payload))
	}
}

export function fetchedEdit(user) {
	return (dispatch) => {
		let payload = {
			loading: false,
			user: {
				...user
			}
		}
		dispatch(fetchUser(payload))
	}
}

export function saveEditUser(savedUser, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/${savedUser.employee_number}`, {
				method: 'PUT',
				body: JSON.stringify(savedUser)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					message.success(body)
					pusher('/admin')
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("err @saveEditUser: ", error)
			})
	}
}