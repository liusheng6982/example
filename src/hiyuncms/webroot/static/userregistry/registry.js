/**
 * 鍔熻兘璇存槑:		杈撳叆楠岃瘉
 * @author:		vivy <lizhizyan@qq.com>
 * @time:		2015-9-25 16:15:30
 * @version:		V1.1.0
 * @浣跨敤鏂规硶:
 * <input class="required" type="text" data-valid="isNonEmpty||isEmail" data-error="email涓嶈兘涓虹┖||閭鏍煎紡涓嶆纭�" id="" />
 * 1銆侀渶瑕侀獙璇佺殑鍏冪礌閮藉姞涓娿€恟equired銆戞牱寮�
 * 2銆丂data-valid		楠岃瘉瑙勫垯锛岄獙璇佸涓鍒欎腑闂寸敤銆恷|銆戦殧寮€锛屾洿澶氶獙璇佽鍒欙紝鐪媟ules鍜宺ule锛屽悗闈㈤亣鍒板彲缁х画澧炲姞
 * 3銆丂data-error		瑙勫垯瀵瑰簲鐨勬彁绀轰俊鎭紝涓€涓€瀵瑰簲
 *
 * @js璋冪敤鏂规硶锛�
 * verifyCheck({
*  	formId:'verifyCheck',		<楠岃瘉formId鍐卌lass涓簉equired鐨勫厓绱�
*	onBlur:null,				<琚獙璇佸厓绱犲け鍘荤劍鐐圭殑鍥炶皟鍑芥暟>
*	onFocus:null,				<琚獙璇佸厓绱犺幏寰楃劍鐐圭殑鍥炶皟鍑芥暟>
*	onChange: null,				<琚獙璇佸厓鍊兼敼鍙樼殑鍥炶皟鍑芥暟>
*	successTip: true,			<楠岃瘉閫氳繃鏄惁鎻愮ず>
*	resultTips:null,			<鏄剧ず鎻愮ず鐨勬柟娉曪紝鍙傛暟obj[褰撳墠鍏冪礌],isRight[鏄惁姝ｇ‘鎻愮ず],value[鎻愮ず淇℃伅]>
*	clearTips:null,				<娓呴櫎鎻愮ず鐨勬柟娉曪紝鍙傛暟obj[褰撳墠鍏冪礌]>
*	code:true					<鏄惁闇€瑕佹墜鏈哄彿鐮佽緭鍏ユ帶鍒堕獙璇佺爜鍙婄偣鍑婚獙璇佺爜鍊掕鏃�,鐩墠鍥哄畾鎵嬫満鍙风爜ID涓簆hone,楠岃瘉鐮佷袱涓爣绛緄d鍒嗗埆涓簍ime_box锛宺esend,濉啓楠岃瘉妗唅d涓篶ode>
*	phone:true					<鏀瑰彉鎵嬫満鍙锋椂鏄惁鎺у埗楠岃瘉鐮�>
* })
 * $("#submit-botton").click(function(){		<鐐瑰嚮鎻愪氦鎸夐挳鏃堕獙璇�>
*  	if(!common.verify.btnClick()) return false;
* })
 *
 * 璇︾粏浠ｇ爜璇风湅register.src.js
 */
(function($) {
    var h, timerC = 60,
        opt;
    var j = function(a) {
        a = $.extend(require.defaults, a || {});
        opt = a;
        return (new require())._init(a)
    };

    function require(f) {
        var g = {
            phone: /^1(3\d|5[0-35-9]|8[025-9]|47)\d{8}$/,
            company: /^[涓€-榫-zA-Z][涓€-榫-zA-Z0-9\s-,-.]*$/,
            uname: /^[涓€-榫-zA-Z][涓€-榫-zA-Z0-9_]*$/,
            zh: /^[涓€-榫+$/,
			card: /^((1[1-5])|(2[1-3])|(3[1-7])|(4[1-6])|(5[0-4])|(6[1-5])|71|(8[12])|91)\d{4}(((((19|20)((\d{2}(0[13-9]|1[012])(0[1-9]|[12]\d|30))|(\d{2}(0[13578]|1[02])31)|(\d{2}02(0[1-9]|1\d|2[0-8]))|(([13579][26]|[2468][048]|0[48])0229)))|20000229)\d{3}(\d|X|x))|(((\d{2}(0[13-9]|1[012])(0[1-9]|[12]\d|30))|(\d{2}(0[13578]|1[02])31)|(\d{2}02(0[1-9]|1\d|2[0-8]))|(([13579][26]|[2468][048]|0[48])0229))\d{3}))$/,
            int: /^[0-9]*$/,
            s: ''
        };
        this.rules = {
            isNonEmpty: function(a, b) {
                b = b || " ";
                if (!a.length) return b
            },
            minLength: function(a, b, c) {
                c = c || " ";
                if (a.length < b) return c
            },
            maxLength: function(a, b, c) {
                c = c || " ";
                if (a.length > b) return c
            },
            isRepeat: function(a, b, c) {
                c = c || " ";
                if (a !== $("#" + b).val()) return c
            },
            between: function(a, b, c) {
                c = c || " ";
                var d = parseInt(b.split('-')[0]);
                var e = parseInt(b.split('-')[1]);
                if (a.length < d || a.length > e) return c
            },
            level: function(a, b, c) {
                c = c || " ";
                var r = j.pwdStrong(a);
                if (b > 4) b = 3;
                if (r < b) return c
            },
            isPhone: function(a, b) {
                b = b || " ";
                if (!g.phone.test(a)) return b
            },
            isCompany: function(a, b) {
                b = b || " ";
                if (!g.company.test(a)) return b
            },
            isInt: function(a, b) {
                b = b || " ";
                if (!g.int.test(a)) return b
            },
            isUname: function(a, b) {
                b = b || " ";
                if (!g.uname.test(a)) return b
            },
            isZh: function(a, b) {
                b = b || " ";
                if (!g.zh.test(a)) return b
            },
            isCard: function(a, b) {
                b = b || " ";
                if (!g.card.test(a)) return b
            },
            isChecked: function(c, d, e) {
                d = d || " ";
                var a = $(e).find('input:checked').length,
                    b = $(e).find('.on').length;
                if (!a && !b) return d
            }
        }
    };
    require.prototype = {
        _init: function(b) {
            this.config = b;
            this.getInputs = $('#' + b.formId).find('.required:visible');
            var c = false;
            var d = this;
            if (b.code) {
                $("#verifyYz").click(function() {
                    $("#time_box").text("60 s鍚庡彲閲嶅彂");
                    d._sendVerify()
                })
            }
            $('body').on({
                blur: function(a) {
                    d.formValidator($(this));
                    if (b.phone && $(this).attr("id") === "phone") d._change($(this));
                    b.onBlur ? b.onBlur($(this)) : ''
                },
                focus: function(a) {
                    b.onFocus ? b.onFocus($(this)) : $(this).parent().find("label.focus").not(".valid").removeClass("hide").siblings(".valid").addClass("hide") && $(this).parent().find(".blank").addClass("hide") && $(this).parent().find(".close").addClass("hide")
                },
                keyup: function(a) {
                    if (b.phone && $(this).attr("id") === "phone") d._change($(this))
                },
                change: function(a) {
                    b.onChange ? b.onChange($(this)) : ''
                }
            }, "#" + b.formId + " .required:visible");
            $('body').on("click", ".close", function() {
                var p = $(this).parent(),
                    input = p.find("input");
                input.val("").focus()
            })
        },
        formValidator: function(a) {
            var b = a.attr('data-valid');
            if (b === undefined) return false;
            var c = b.split('||');
            var d = a.attr('data-error');
            if (d === undefined) d = "";
            var e = d.split("||");
            var f = [];
            for (var i = 0; i < c.length; i++) {
                f.push({
                    strategy: c[i],
                    errorMsg: e[i]
                })
            };
            return this._add(a, f)
        },
        _add: function(a, b) {
            var d = this;
            for (var i = 0, rule; rule = b[i++];) {
                var e = rule.strategy.split(':');
                var f = rule.errorMsg;
                var g = e.shift();
                e.unshift(a.val());
                e.push(f);
                e.push(a);
                var c = d.rules[g].apply(a, e);
                if (c) {
                    opt.resultTips ? opt.resultTips(a, false, c) : j._resultTips(a, false, c);
                    return false
                }
            }
            opt.successTip ? (opt.resultTips ? opt.resultTips(a, true) : j._resultTips(a, true)) : j._clearTips(a);
            return true
        },
        _sendVerify: function() {
            var a = this;
            $("#verifyYz").text("鍙戦€侀獙璇佺爜").hide();
            $("#time_box").text("10 s鍚庡彲閲嶅彂").show();
            if (timerC === 0) {
                clearTimeout(h);
                timerC = 60;
                var b = /^1([^01269])\d{9}$/;
                if (!b.test($("#phone").val())) {
                    $("#time_box").text("鍙戦€侀獙璇佺爜")
                } else {
                    $("#verifyYz").show();
                    $("#time_box").hide()
                }
                return
            }
            $("#time_box").text(timerC + " s鍚庡彲閲嶅彂");
            timerC--;
            h = setTimeout(function() {
                a._sendVerify()
            }, 1000)
        },
        _change: function(a) {
            var b = this;
            if (a.val().length != 11) {
                $("#verifyYz").hide();
                $("#time_box").show();
                if (timerC === 60) $("#time_box").text("鍙戦€侀獙璇佺爜");
                $("#verifyNo").val("");
                this.config.clearTips ? this.config.clearTips($("#verifyNo")) : j._clearTips($("#verifyNo"));
                return
            }
            var c = /^1([^01269])\d{9}$/;
            if (!c.test(a.val())) return false;
            if (timerC === 60) {
                $("#verifyYz").show();
                $("#time_box").hide()
            } else {
                $("#verifyYz").hide();
                $("#time_box").show()
            }
        }
    };
    j._click = function(c) {
        c = c || opt.formId;
        var d = $("#" + c).find('.required:visible'),
            self = this,
            result = true,
            t = new require(),
            r = [];
        $.each(d, function(a, b) {
            result = t.formValidator($(b));
            if (result) r.push(result)
        });
        if (d.length !== r.length) result = false;
        return result
    };
    j._clearTips = function(a) {
        a.parent().find(".blank").addClass("hide");
        a.parent().find(".valid").addClass("hide");
        a.removeClass("v_error")
    };
    j._resultTips = function(a, b, c) {
        a.parent().find("label.focus").not(".valid").addClass("hide").siblings(".focus").removeClass("hide");
        a.parent().find(".close").addClass("hide");
        a.removeClass("v_error");
        c = c || "";
        if (c.length > 21) c = "<span>" + c + "</span>";
        var o = a.parent().find("label.valid");
        if (!b) {
            o.addClass("error");
            a.addClass("v_error");
            if ($.trim(a.val()).length > 0) a.parent().find(".close").removeClass("hide")
        } else {
            a.parent().find(".blank").removeClass("hide")
        }
        o.text("").append(c)
    };
    j.textChineseLength = function(a) {
        var b = /[涓€-榫|[銆�-銆俔|[锛�-锛焆|[銆�-銆廬|[銆�-銆昡|[鈥�-鈥漖|[锛�-锛嶿|[銆�-銆塢|[鈥|[锟/g;
		if (b.test(a)) return a.match(b).length;
		else return 0
	};
	j.pwdStrong = function(a) {
		var b = 0;
		if (a.match(/[a-z]/g)) {
    b++
}
if (a.match(/[A-Z]/g)) {
    b++
}
if (a.match(/[0-9]/g)) {
    b++
}
if (a.match(/(.[^a-z0-9A-Z])/g)) {
    b++
}
if (b > 4) {
    b = 4
}
if (b === 0) return false;
return b
};
require.defaults = {
    formId: 'verifyCheck',
    onBlur: null,
    onFocus: null,
    onChange: null,
    successTip: true,
    resultTips: null,
    clearTips: null,
    code: true,
    phone: false
};
window.verifyCheck = $.verifyCheck = j
})(jQuery);
(function($) {
    var f;
    var g = function() {
        return (new require())._init()
    };

    function require(a) {};
    require.prototype = {
        _init: function() {
            var b = this;
            $('body').on({
                click: function(a) {
                    b._click($(this))
                }
            }, ".showpwd:visible")
        },
        _click: function(a) {
            var c = a.attr('data-eye');
            if (c === undefined) return false;
            var d = $("#" + c),
                cls = !d.attr("class") ? "" : d.attr("class"),
                value = !d.val() ? "" : d.val(),
                type = d.attr("type") === "password" ? "text" : "password",
                b = d.parent().find("b.placeTextB"),
                isB = b.length === 0 ? false : true;
            var s = d.attr("name") ? " name='" + d.attr("name") + "'" : "";
            s += d.attr("data-valid") ? " data-valid='" + d.attr("data-valid") + "'" : "";
            s += d.attr("data-error") ? " data-error='" + d.attr("data-error") + "'" : "";
            s += d.attr("placeholder") ? " placeholder='" + d.attr("placeholder") + "'" : "";
            var e = '<input readonly type="' + type + '" class="' + cls + '" value="' + value + '" id="' + c + '"' + s + ' />';
            if (type === "text") {
                if (isB) b.hide();
                d.parent().find(".icon-close.close").addClass("hide");
                d.removeAttr("id").hide();
                d.after(e);
                a.addClass("hidepwd")
            } else {
                d.prev("input").attr("id", c).val(value).show();
                if (isB && $.trim(value) === "") {
                    d.prev("input").hide();
                    b.show()
                }
                d.remove();
                a.removeClass("hidepwd")
            };
            $('body').on("click", "#" + c, function() {
                $(this).parent().find(".hidepwd").click();
                if (isB && $.trim($(this).val()) === "") {
                    d.show();
                    b.hide()
                }
                d.focus()
            })
        }
    };
    require.defaults = {};
    window.togglePwd = $.togglePwd = g
})(jQuery);
(function($) {
    var b, timerC, opt;
    var d = function(a) {
        a = $.extend(require.defaults, a || {});
        opt = a;
        d._clear();
        return (new require())._init()
    };

    function require(a) {};
    require.prototype = {
        _init: function() {
            timerC = opt.maxTime;
            this._sendVerify()
        },
        _sendVerify: function() {
            var a = this;
            if (timerC === 0) {
                d._clear();
                opt.after();
                timerC = opt.maxTime;
                return
            }
            timerC--;
            opt.ing(timerC);
            b = setTimeout(function() {
                a._sendVerify()
            }, 1000)
        }
    };
    d._clear = function() {
        clearTimeout(b)
    };
    require.defaults = {
        maxTime: 60,
        minTime: 0,
        ing: function(c) {},
        after: function() {}
    };
    window.countdown = $.countdown = d
})(jQuery);
$(function() {
    togglePwd();
    verifyCheck();
    $('body').on("keyup", "#password", function() {
        var t = $(this).val(),
            o = $(this).parent().find(".strength");
        if (t.length >= 6) {
            o.show();
            var l = verifyCheck.pwdStrong(t);
            o.find("b i").removeClass("on");
            for (var i = 0; i < l; i++) {
                o.find("b i").eq(i).addClass("on")
            }
        } else {
            o.hide()
        }
    })
});