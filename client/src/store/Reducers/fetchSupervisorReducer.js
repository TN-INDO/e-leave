import {
    FETCH_LEAVE_PENDING,
    FETCH_LEAVE_APPROVE,
    FETCH_LEAVE_REJECT,
    APPROVE_LEAVE_PENDING,
    REJECT_LEAVE_PENDING
} from "../Actions/types"

let supervisorState = {
    loading: true,
    leaves: [],
    leave: {
        reject_reason: ''
    }
}

export default function fetchSupervisorReducer(state = supervisorState, action) {
    switch (action.type) {
        case FETCH_LEAVE_PENDING:
            return {
                ...action.payload
            }
        case FETCH_LEAVE_APPROVE:
            return {
                ...action.payload
            }
        case FETCH_LEAVE_REJECT:
            return {
                ...action.payload
            }
        case APPROVE_LEAVE_PENDING:
            return {
                ...action.payload
            }
        case REJECT_LEAVE_PENDING:
            return {
                ...action.payload
            }
        case 'REJECT_REASON':
            return {
                ...action.payload
            }
        default:
            return state
    }
}