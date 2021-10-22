const isRelease = false;

export default {
    isRelease,
    frontend: isRelease ? "" : "http://localhost:8100",
    backend: isRelease ? "" : "http://localhost:8080"
};
