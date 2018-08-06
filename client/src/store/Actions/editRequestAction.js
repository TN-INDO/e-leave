import {
	ROOT_API
} from "./types.js";
import {
	message
} from "antd";

function fetchRequest(payload) {
	return {
		type: 'FETCH_EDIT',
		payload: payload
	}
}

function requestEditing(payload) {
	return {
		type: 'EDITING_REQUEST',
		payload: payload
	}
}

export function handleEdit(newLeave) {
	return (dispatch) => {
		let payload = {
			loading: false,
			leave: {
				...newLeave
			}
		}
		dispatch(requestEditing(payload))
	}
}

export function fetchedEdit(leave) {
	return (dispatch) => {
		let payload = {
			loading: false,
			leave: {
				...leave
			}
		}
		dispatch(fetchRequest(payload))
	}
}

export function saveEditLeave(savedLeave, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/leave/${savedLeave.id}`, {
				method: 'PUT',
				body: JSON.stringify(savedLeave)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					message.success(body)
					pusher('/employee/list-pending-request')
				} else if (error !== null) {
					message.error(error)
				}
			}).catch(error => {
				console.error(error)
			})
	}
}