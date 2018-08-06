import {
	ROOT_API,
	FETCH_REQUEST_PENDING,
	DELETE_REQUEST_PENDING,
	FETCH_REQUEST_APPROVE,
	FETCH_REQUEST_REJECT,
} from "./types"

function fetchRequestPending(payload) {
	return {
		type: FETCH_REQUEST_PENDING,
		payload: payload
	}
}

function deleteRequestPending(payload) {
	return {
		type: DELETE_REQUEST_PENDING,
		payload: payload
	}
}

function fetchRequestApprove(payload) {
	return {
		type: FETCH_REQUEST_APPROVE,
		payload: payload
	}
}

function fetchRequestReject(payload) {
	return {
		type: FETCH_REQUEST_REJECT,
		payload: payload
	}
}

export function employeeGetRequestPending() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/pending/${employeeNumber}`, {
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
				dispatch(fetchRequestPending(payload))

				if (error !== null) {
					console.error("error not null @employeeGetRequestPending: ", error)
				}
			})
			.catch(error => {
				console.error("error @employeeGetRequestPending: ", error)
			})
	}
}

export function employeeDeleteRequest(leaves, id) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/leave/${id}`, {
				method: 'DELETE',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let newUserlist = leaves.filter(el => el.id !== id)
				let payload = {
					loading: false,
					leaves: [
						...newUserlist
					]
				}
				dispatch(deleteRequestPending(payload))

				if (error !== null) {
					console.error("error not null @employeeDeleteRequest: ", error)
				}
			})
			.catch(error => {
				console.error("error @employeeDeleteRequest: ", error)
			})
	}
}

export function employeeGetRequestApprove() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/accept/${employeeNumber}`, {
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
				dispatch(fetchRequestApprove(payload))

				if (error !== null) {
					console.error("error not null @employeeGetRequestApprove: ", error)
				}
			})
			.catch(error => {
				console.error("error @employeeGetRequestApprove: ", error)
			})
	}
}

export function employeeGetRequestReject() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/employee/reject/${employeeNumber}`, {
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
				dispatch(fetchRequestReject(payload))

				if (error !== null) {
					console.error("error not null @employeeGetRequestReject: ", error)
				}
			})
			.catch(error => {
				console.error("error @employeeGetRequestReject: ", error)
			})
	}
}