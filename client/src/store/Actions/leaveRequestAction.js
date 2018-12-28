import {
	ROOT_API,
	CREATE_LEAVE_REQUEST,
	CLEAR_FIELD
} from "./types"
import {
	message
} from "antd";

export function formOnChange(payload) {
	return (dispach) => {
		dispach({
			type: CREATE_LEAVE_REQUEST,
			payload: payload
		})
	}
}

function clearField() {
	return {
		type: CLEAR_FIELD
	}
}

export function SumbitLeave(payload, pusher) {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/leave/${employeeNumber}`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {				
				if (body !== null) {
					dispatch(clearField())
					pusher('/employee/list-pending-request')
					message.success('Create leave request success')
				} else if (error === "Type request malform") {
					message.error('Create failed, please check all field!')
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @SumbitLeave: ", error)
			})
	}
}

export function SumbitLeaveSupervisor(payload, pusher) {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/leave/${employeeNumber}`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					dispatch(clearField())
					pusher('/employee/list-pending-request')
					message.success('Create leave request success')
				} else if (error === "Type request malform") {
					message.error('Create failed, please check all field!')
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @SumbitLeaveSupervisor: ", error)
			})
	}
}