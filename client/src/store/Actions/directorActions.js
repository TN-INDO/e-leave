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

export function fetchDirectorLeavePending() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/director/pending`, {
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
					console.error("error not null @fetchDirectorLeavePending: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchDirectorLeavePending: ", error)
			})
	}
}

export function fetchDirectorLeaveApprove() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/director/accept`, {
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
					console.error("error not null @fetchDirectorLeaveApprove: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchDirectorLeaveApprove: ", error)
			})
	}
}

export function fetchDirectorLeaveReject() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/director/reject/`, {
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
					console.error("error not null @fetchDirectorLeaveReject: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchDirectorLeaveReject: ", error)
			})
	}
}

export function directorApproveRequest(leave, id, enumber) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/director/accept/${id}/${enumber}/`, {
				method: 'PUT',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error,
			}) => {
				if (body !== null) {
					let newLeavelist = leave.filter(el => el.id !== id)
					let payload = {
						loading: false,
						leaves: [
							...newLeavelist
						]
					}
					dispatch(approveRequest(payload))
					message.success(body)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @directorApproveRequest: ", error)
			})
	}
}

export function directorRejectRequest(leave, id, enumber, payload) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/director/reject/${id}/${enumber}/`, {
				method: 'PUT',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					let newLeavelist = leave.filter(el => el.id !== id)
					let payloads = {
						loading: false,
						leaves: [
							...newLeavelist
						]
					}
					dispatch(rejectRequest(payloads))
					message.success(body)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @directorRejectRequest: ", error)
			})
	}
}