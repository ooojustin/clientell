import { createStore } from "vuex";
//import { Http } from "@capacitor-community/http";

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
        },
        async restoreLogin({ commit }) {
            const token = localStorage.getItem("token");
            console.log("restore login", token);
        }
    },

    mutations: {
        loggedIn(state: any, payload: any) {
            state.user = payload;
            state.token = payload.token;
            localStorage.setItem("token", state.token);
        }
    },

    getters: {
        isAuthenticated(state: any) {
            return state.token.length > 0;
        }
    }

});
