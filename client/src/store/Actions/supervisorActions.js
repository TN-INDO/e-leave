import {
	ROOT_API,
	FETCH_LEAVE_PENDING,
	FETCH_LEAVE_APPROVE,
	FETCH_LEAVE_REJECT,
	APPROVE_LEAVE_PENDING,
	REJECT_LEAVE_PENDING
} from "./types"
import {
	message
} from "antd";

function pendingFetch(payload) {
	return {
		type: FETCH_LEAVE_PENDING,
		payload: payload
	}
}

function approveFetch(payload) {
	return {
		type: FETCH_LEAVE_APPROVE,
		payload: payload
	}
}

function rejectFetch(payload) {
	return {
		type: FETCH_LEAVE_REJECT,
		payload: payload
	}
}

function approveRequest(payload) {
	return {
		type: APPROVE_LEAVE_PENDING,
		payload: payload
	}
}

function rejectRequest(payload) {
	return {
		type: REJECT_LEAVE_PENDING,
		payload: payload
	}
}

export function fetchSupervisorLeavePending() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/pending/${employeeNumber}`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					leaves: body
				}
				dispatch(pendingFetch(payload))

				if (error !== null) {
					console.error("error not null @fetchSupervisorLeavePending: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchSupervisorLeavePending: ", error)
			})
	}
}

export function fetchSupervisorLeaveApprove() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/accept/${employeeNumber}`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					leaves: body
				}
				dispatch(approveFetch(payload))

				if (error !== null) {
					console.error("error not null @fetchSupervisorLeaveApprove: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchSupervisorLeaveApprove: ", error)
			})
	}
}

export function fetchSupervisorLeaveReject() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/reject/${employeeNumber}`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					leaves: body
				}
				dispatch(rejectFetch(payload))

				if (error !== null) {
					console.error("error not null @fetchSupervisorLeaveReject: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchSupervisorLeaveReject: ", error)
			})
	}
}

export function supervisorApproveRequest(leaves, id, enumber) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/accept/${id}/${enumber}`, {
				method: 'PUT',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					let listNewLeave = leaves.filter(el => el.id !== id)
					let payload = {
						loading: false,
						leaves: [
							...listNewLeave
						]
					}
					dispatch(approveRequest(payload))
					message.success(body)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @supervisorApproveRequest: ", error)
			})
	}
}

export function supervisorRejectRequest(leaves, id, enumber, payload) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/supervisor/reject/${id}/${enumber}`, {
				method: 'PUT',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					let listNewLeave = leaves.filter(el => el.id !== id)
					let payloads = {
						loading: false,
						leaves: [
							...listNewLeave
						]
					}
					dispatch(rejectRequest(payloads))
					message.success(body)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @supervisorRejectRequest: ", error)
			})
	}
}