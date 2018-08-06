import {
	ROOT_API,
	FETCH_TYPE_LEAVE
} from "./types"

function typeLeaveFetch(payload) {
	return {
		type: FETCH_TYPE_LEAVE,
		payload: payload
	}
}

export function typeLeaveFetchData() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/user/type-leave/`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					typeLeave: body
				}
				dispatch(typeLeaveFetch(payload))

				if (error !== null) {
					console.error("error not null @typeLeaveFetchData: ", error)
				}
			})
			.catch(error => {
				console.error("error @typeLeaveFetchData: ", error)
			})
	}
}