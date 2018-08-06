import {
	ROOT_API,
	FETCH_USER,
	DELETE_USER,
	FETCH_LEAVE_PENDING,
	FETCH_LEAVE_APPROVE,
	FETCH_LEAVE_REJECT,
	CANCEL_LEAVE_REQUEST,
	FETCH_LEAVE_BALANCES
} from "./types"
import {
	message
} from "antd"

function userFetch(payload) {
	return {
		type: FETCH_USER,
		payload: payload
	}
}

function userDeleted(payload) {
	return {
		type: DELETE_USER,
		payload: payload
	}
}

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

function cancelRequest(payload) {
	return {
		type: CANCEL_LEAVE_REQUEST,
		payload: payload
	}
}

function fetchLeaveBalance(payload) {
	return {
		type: FETCH_LEAVE_BALANCES,
		payload: payload
	}
}

export function adminGetUsers() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					users: body
				}
				dispatch(userFetch(payload))

				if (error !== null) {
					console.error("error not null @adminGetUsers: ", error)
				}
			})
			.catch(error => {
				console.error("error @adminGetUsers: ", error)
			})
	}
}

export function adminEditLeaveBalances(employeeNumber) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/user/type-leave/${employeeNumber}`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					balances: body
				}
				dispatch(fetchLeaveBalance(payload))

				if (error !== null) {
					console.error("error not null @adminEditLeaveBalances: ", error)
				}
			})
			.catch(error => {
				console.error("error @adminEditLeaveBalances: ", error)
			})
	}
}

export function adminDeleteUser(users, employeeNumber) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/${employeeNumber}`, {
				method: 'DELETE',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let newUserlist = users.filter(el => el.employee_number !== employeeNumber)
				let payload = {
					loading: false,
					users: [
						...newUserlist
					]
				}
				dispatch(userDeleted(payload))

				if (error !== null) {
					console.error("error not null @adminDeleteUser: ", error)
				}
			})
			.catch(error => {
				console.error("error @adminDeleteUser: ", error)
			})
	}
}

export function fetchAdminLeavePending() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/leave/pending/`, {
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
					console.error("error not null @fetchAdminLeavePending: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchAdminLeavePending: ", error)
			})
	}
}

export function fetchAdminLeaveApprove() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/leave/accept/`, {
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
					console.error("error not null @fetchAdminLeaveApprove: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchAdminLeaveApprove: ", error)
			})
	}
}


export function fetchAdminLeaveReject() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/leave/reject/`, {
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
					console.error("error not null @fetchAdminLeaveReject: ", error)
				}
			})
			.catch(error => {
				console.error("error @fetchAdminLeaveReject: ", error)
			})
	}
}

export function cancelRequestLeave(users, id, enumber) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/leave/cancel/${id}/${enumber}/`, {
				method: 'PUT',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					let newUserlist = users.filter(el => el.id !== id)
					let payload = {
						loading: false,
						leaves: [
							...newUserlist
						]
					}
					dispatch(cancelRequest(payload))
					message.success(body)
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @cancelRequestLeave: ", error)
			})
	}
}

export function downloadReport(from, to) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/leave/reports?fromDate=${from}&toDate=${to}`, {
				method: 'GET'
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					const url = window.URL.createObjectURL(new Blob([arrayToCSV(body)]));
					const link = document.createElement('a');
					link.href = url;
					link.setAttribute('download', 'report_leave_request.xlsx');
					document.body.appendChild(link);
					link.click();
					message.success('Download success')
					window.location.reload();
				} else if (body === null) {
					message.error('Data is not available')
				} else {
					message.error(error)
				}
			})
			.catch(error => {
				console.error("error @downloadReport: ", error)
			})
	}
}

export function downloadReportTypeLeave(from, to, id) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/leave/report/type?fromDate=${from}&toDate=${to}&typeID=${id}`, {
				method: 'GET'
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					const url = window.URL.createObjectURL(new Blob([arrayToCSV(body)]));
					const link = document.createElement('a');
					link.href = url;
					link.setAttribute('download', 'report_leave_request_by_type_leave.xlsx');
					document.body.appendChild(link);
					link.click();
					message.success('Download success')
					window.location.reload();
				} else if (body === null) {
					message.error('Data is not available')
				} else {
					message.error(error)
				}
			})
			.catch(error => {
				console.error("error @downloadReportTypeLeave: ", error)
			})
	}
}

function arrayToCSV(objArray) {
	const array = typeof objArray !== 'object' ? JSON.parse(objArray) : objArray;
	let str = `${Object.keys(array[0]).map(value => `"${value}"`).join(",")}` + '\r\n'; // eslint-disable-line

	return array.reduce((str, next) => {
		str += `${Object.values(next).map(value => `"${value}"`).join(",")}` + '\r\n'; // eslint-disable-line
		return str;
	}, str);
}