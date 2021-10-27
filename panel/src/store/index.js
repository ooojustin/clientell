import { createStore } from 'vuex';
import vars, { initApi } from '../variables';
import axios from "axios";

export default createStore({
    state: {
        user: null,
        token: ""
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
            const response = await axios.get(`${vars.backend}/user`, { 
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
        loggedIn(state, payload) {
            state.user = payload;
            state.token = payload.token;
            localStorage.setItem("token", state.token);
            initApi(state.token);
        }
    },
    modules: {
    },
    getters: {
        isAuthenticated(state) {
            // whether or not a token is set in state
            return state.token.length > 0;
        }
    }
});
