jQuery(document).ready(function($) {
    let scaleSize = window.innerWidth /1920;
    if (this.scaleSelected)
        document.body.style.WebkitTransform = "scale(" + scaleSize * (2 / 3) + "," + scaleSize + ")";
    else
        document.body.style.WebkitTransform = "scale(" + scaleSize + ")";

})

