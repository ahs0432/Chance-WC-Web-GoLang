function characterCheck(obj) {
    var regExp = /[ \{\}\[\]\/?.,;:|\)*~`!^\-_+┼<>@\#$%&\'\"\\\(\=]/gi; 

    if (regExp.test(obj.value)) {
        alert("특수문자는 입력하실 수 없습니다.");
        obj.value = obj.value.substring(0, obj.value.length - 1);
    }
}