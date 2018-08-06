// import jwtDecode from 'jwt-decode';
import {
	ROOT_API,
	LOGIN_FORM_ONCHANGE,
	CLEAR_FIELD
} from "./types"

import {
	message
} from "antd";


export function handleFormInput(payload) {
	return {
		type: LOGIN_FORM_ONCHANGE,
		payload: payload
	}
}

function clearField() {
	return {
		type: CLEAR_FIELD
	}
}

export function submitLogin(payload, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/login`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (error !== null) {
					message.error(error);
				} else {
					const token = body['Token']
					const id = body['ID']
					const role = body['Role']
										
					if (role === 'admin') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/admin')
						dispatch(clearField())
						message.success('Login success!')
					} else if (role === 'director') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/director')
						dispatch(clearField())
						message.success('Login success!')
					} else if (role === 'supervisor') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/supervisor')
						dispatch(clearField())
						message.success('Login success!')
					} else if (role === 'employee') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/employee')
						dispatch(clearField())
						message.success('Login success!')
					} else if (role !== 'admin' || role !== 'director' || role !== 'supervisor' || role !== 'employee') {
						pusher('/')
						dispatch(clearField())
						message.error('Login failed!');
					}
				}
			})
			.catch(error => {
				console.error("error @submitLogin: ", error)
			})
	}
}