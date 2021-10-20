import { createStore } from "vuex";

export default createStore({

    state() {
        return {
            user: null,
            token: ""
        };
    },

    actions: {
        loggedIn({ commit }, payload) {
            commit("loggedIn", payload);
        }
    },

    mutations: {
        loggedIn(state: any, payload: any) {
            state.user = payload;
            state.token = payload.token;
        }
    },

    getters: {
        isAuthenticated(state: any) {
            return state.token.length > 0;
        }
    }

});
