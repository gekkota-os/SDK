(function(global) {
	"use strict";
	if("AppLogger" in global){
		var logger =  new AppLogger("app.ontologies.XfpDetectionHid");
	}
	else{
		var logger = new XpfLogger("app.ontologies.XpfDetectionHid");
	}

	function callStateFunction(post, state) {
		var funName = "onState" + state + post;
		logger.debug(funName);
		if(funName in result){
			var fun = result[funName];
			if(fun instanceof Function) {
				logger.debug(fun);
				try {
					fun();
				} catch(e) {
					logger.errorEx(e);
				}
			}
		}
	}


	var isIdle = undefined;

	function setIdle(aIsIdle){
		function getPost(aIsIdle){
			if(aIsIdle){
				return "idle";
			}
			return "active";
		}

		if(aIsIdle === isIdle){
			return;
		}
		if(isIdle !== undefined){
			callStateFunction("End", getPost(isIdle));
		}
		isIdle = aIsIdle;
		if(isIdle !== undefined){
			callStateFunction("Begin", getPost(isIdle));
		}
	}

	var obs = {
		time: 10, 
		onidle: function(){
			logger.debug("onidle callback");
			setIdle(true);
		},
		onactive: function(){
			logger.debug("onactive callback");
			setIdle(false);
		}
	};

	var obsAttached = false;
	function attachObserver(){
		logger.debug("attach observer");
		try{
			navigator.addIdleObserver(obs);
			obsAttached = true;
		}catch(e){
			logger.errorEx(e);
		}
	}

	function detachObserver(){
		logger.debug("detach observer");
		try{
			navigator.removeIdleObserver(obs);
			obsAttached = false;
		}catch(e){
			logger.errorEx(e);
		}
	}

	function unload(){
		setIdle(undefined);
		if(obsAttached){
			detachObserver();
		}
	}

	function load(){
		if(!obsAttached){
			attachObserver();
		}
		if(isIdle === undefined){
			setIdle(false);
		}
		global.removeEventListener("load", load, false);
	}

	global.addEventListener("load", load, false);
	global.addEventListener("unload", unload, false);

	var result = global.XpfDetectionHid = {
		init : function init(val){
			logger.debug("init :" + val);
			if(val == obs.time){
				return;
			}
			if(obsAttached){
				detachObserver();
			}
			obs.time = val;
			attachObserver();
		}
	};
})(window)