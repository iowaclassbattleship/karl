window.onload = () => {
    setTimeout(() => {
        document.getElementById("target").scrollIntoView({ behavior: "smooth" });
        setTimeout(() => {
            document.location = "/";
        }, 1000);
    }, 2500);
};