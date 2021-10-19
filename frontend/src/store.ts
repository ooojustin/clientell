import { createStore } from "vuex";

export const store = createStore({

    state() {
        return {
            token: ""
        };
    },

    actions: {
        setToken({ commit }, payload) {
            commit("setToken", payload);
        }
    },

    mutations: {
        setToken(state: any, payload: { token: string }) {
            state.token = payload.token;
            console.log("set token to:", state.token);
        }
    }

});
