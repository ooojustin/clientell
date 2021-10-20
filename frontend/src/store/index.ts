import { createStore } from "vuex";
import { Http } from "@capacitor-community/http";
import vars from "../variables";

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

            // look for token in local storage
            const token = localStorage.getItem("token");
            if (!token)
                return;

            // send request to get user data from token
            const response = await Http.get({
                url: `${vars.backend}/user`,
                headers: { Token: token }
            });

            // if request succeeded, log user back in
            if (response.status == 200) {
                const user = response.data.data;
                commit("loggedIn", user);
            }
            
        }
    },

    mutations: {
        loggedIn(state: any, payload: any) {
            // load user/token into state, store token in localStorage for future use
            state.user = payload;
            state.token = payload.token;
            localStorage.setItem("token", state.token);
        }
    },

    getters: {
        isAuthenticated(state: any) {
            // whether or not a token is set in state
            return state.token.length > 0;
        }
    }

});
