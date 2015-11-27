# Détection HID

## Présentation 

Dans le cadre de la détection HID, on considère que le player peut être soit dans l'état *Actif* soit dans l'état *Inactif*.
Le passage d'un état à un autre se base sur les évènements clavier/souris/écran (événements HID avec focus sur le player) et sur un délai fixe (valeur par défaut de 5 secondes ou spécifiée par l'utilisateur).

Dès qu'un évènement HID arrive, le player passe ou reste en mode *Actif*. A son démarrage, le player est dans l'état *Actif*.  
S'il ne se passe aucun évènement HID entre le moment du dernier évènement HID et la fin du délai, le player passe en mode *Inactif*.

## Fonction d'initialisation
- Initialisation : cette fonction prend en paramètre un entier (*unsigned long*) minimal de 1. Il s'agit du délai en secondes avant que le player ne soit considéré en mode *Inactif*.  
On peut placer cette fonction n'importe quand. Si la détection est déjà en route, la fonction aura seulement pour effet de mettre à jour le délai.  
Si cette fonction n'est jamais appelée, la détection est initialisée par défaut avec un délai de 5 secondes.

Exemple -> placer *Initialisation(30)* dans le calendrier au moment du *Début* de la journée. Dans ce cas, à partir du début de la journée, le player passera en état *Inactif* uniquement quand il n'aura pas reçu d'évènement HID pendant 30 secondes.


## Variables

### Variable *Actif* 

- début : *exemple -> jouer un message de bienvenue*
- pendant : *exemple -> jouer une séquence de médias*
- fin

Si l'on place une séquence sur *Pendant* , elle sera forcément jouée lors de l'état *Actif*, même s'il y a une séquence configurée dans le calendrier.  
Si l'on ne place rien sur *Pendant*, c'est la séquence configurée sur le calendrier qui sera jouée lors de l'état *Actif*.

### Variable *Inactif* 

- début : *exemple -> baisser l'éclairage de l'écran*
- pendant : 
- fin : *exemple -> augmenter l'éclairage de l'écran* 

Si l'on place une séquence sur *Pendant* , elle sera forcément jouée lors de l'état *Inactif*, même s'il y a une séquence configurée dans le calendrier.  
Si l'on ne place rien sur *Pendant*, c'est la séquence configurée sur le calendrier qui sera jouée lors de l'état *Inactif*.

## Exemple de code Gekkota utilisant xpfDetectionHid

	/** @brief example of js code generated for Gekkota */

	/** Start a media at the beginning of the Active state*/
	XpfDetectionHid.onStateactiveBegin = function onStateactiveBegin() {
		try {
			xpfLogger.debug("XpfDetectionHid.onStateactiveBegin()");
			XpfTimingManager.addOverrideElement(document.getElementById("UIDde574046_1824_4dc6_9403_ce895f59f42e"), false, true);
		} catch(e){
			xpfLogger.errorEx(e);
		}
	}

	/** Close a media at the end of the Active state*/
	XpfDetectionHid.onStateactiveEnd = function onStateactiveEnd() {
		try {
			xpfLogger.debug("XpfDetectionHid.onStateactiveEnd()");
			XpfTimingManager.removeOverrideElement(document.getElementById("UIDde574046_1824_4dc6_9403_ce895f59f42e"));
		} catch(e){
			xpfLogger.errorEx(e);
		}
	}

	/** Decrease the brightness at the beginning of the Inactive state*/
	XpfDetectionHid.onStateidleBegin = function onStateidleBegin() {
		try {
			xpfLogger.debug("XpfDetectionHid.onStateidleBegin()");
			try {
				xpfLogger.debug("calling XpfDisplay.brightness(40);");
				XpfDisplay.brightness(40);
			} catch(e){
				xpfLogger.errorEx(e);
			}
			XpfTimingManager.addOverrideElement(document.getElementById("UID88ee92db_e02c_4a30_8b4d_c579b5b02e35"), false, true);
		} catch(e){
			xpfLogger.errorEx(e);
		}
	}

	/** Increase the brightness at the end of the Inactive state*/
	XpfDetectionHid.onStateidleEnd = function onStateidleEnd() {
		try {
			xpfLogger.debug("XpfDetectionHid.onStateidleEnd()");
			try {
				xpfLogger.debug("calling XpfDisplay.brightness(100);");
				XpfDisplay.brightness(100);
			} catch(e){
				xpfLogger.errorEx(e);
			}
			XpfTimingManager.removeOverrideElement(document.getElementById("UID88ee92db_e02c_4a30_8b4d_c579b5b02e35"));
		} catch(e){
			xpfLogger.errorEx(e);
		}
	}
	/** Set the time to inactivity to 7 seconds */
	function onUID4899b153_a0ae_4430_9b5c_4e49f7464923Begin() {
		try {
			xpfLogger.debug("Begin of event 06/11/15 07:36 - 18:18");
			try {
				xpfLogger.debug("calling XpfDetectionHid.init(07);");
				XpfDetectionHid.init(07);
			} catch(e){
				xpfLogger.errorEx(e);
			}
		} catch(e){
			xpfLogger.errorEx(e);
		}
	}