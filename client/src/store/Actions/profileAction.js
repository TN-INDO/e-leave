import {
	ROOT_API,
	GET_PROFILE
} from "./types"

function getProfile(payload) {
	return {
		type: GET_PROFILE,
		payload: payload
	}
}

export function profileFetchData() {
	const employeeNumber = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/${employeeNumber}/`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					user: body
				}
				dispatch(getProfile(payload))

				if (error !== null) {
					console.error("error not null @profileFetchData: ", error)
				}
			})
			.catch(error => {
				console.error("error @profileFetchData: ", error)
			})
	}
}