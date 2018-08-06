import {
    GET_SUPERVISORS
} from "../Actions/types"

let supervisorState = {
    supervisor: [],
}

export default function AddSupervisorReducer(state = supervisorState, action) {
    switch (action.type) {
        case GET_SUPERVISORS:
            return {
                ...action.payload
            }
        default:
            return state
    }
}