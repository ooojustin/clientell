const isRelease = false;

export default {
    isRelease,
    backend: isRelease ? "https://clientellapp.com:8000" : "http://localhost:8080"
};
