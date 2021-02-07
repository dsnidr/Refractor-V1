import { SET_CURRENT_USER } from "./constants";

const initialState = {
    self: null,
}

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case SET_CURRENT_USER:
            return {
                ...state,
                self: action.payload,
            }
        default:
            return state;
    }
}

export default reducer;