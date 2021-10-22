const isRelease = true;

export default {
    isRelease,
    backend: isRelease ? "https://rc.justin.ooo:8000" : "http://localhost:8080"
};
