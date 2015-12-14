# App  xpf-example1

This App displays an image "yellow.jpg" during 3 seconds then an image "example.svg" indefinitly.

## Description of the XPF file content

XPF version 3  is a XML-based format. The XML header  appears at the top of the document :

`<?xml version="1.0" encoding="UTF-8"?>`



## Element <*XPF*>
This root element defines the XML and XPF namespaces.  

`xmlns="ns.innes.xpf.3"`
`xmlns:xpf="ns.innes.xpf.3"`

It is the parent of  <*HEAD*> and <*BODY*> elements.

###  Element <*HEAD*>
This element is the parent of the <*STYLE*> element.

#### Element <*STYLE*>
This element defines the style for : 
 
* general xpf: 

```html  
xpf {		 	
	background-color:#000000;
} 
``` ;  
* body:

```html  
body { 
	overflow:hidden; 
}
``` ;
* and regions:

```html
[region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63"] { 
	top:0px;  
	left:0px;  
	width:100%;  
	height:100%;  
	z-index:0;  
	opacity:1;  
}
```



### Element <*BODY*>
This element contains the PAR element.

#### Element <*PAR*> 
The <*PAR*> element plays child elements as a group.

``` 
<par dur="indefinite">
``` 

##### Attribute *@dur*
This attribute defines the duration for playing the group of childs.  
"indefinite" means there is no time limit.
If the duration was "3s", then the group would stop after 3 seconds.  
##### Element <*SEQ*>
This element defines the medias to play and when to play them.

```
<seq begin="0s" region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" >
	<img class="visual" dur="00:00:03.000" 
	region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" src="medias/yellow.jpg"/>
</seq>
```
```
<seq begin="3s" region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" repeatCount="indefinite">
	<img class="visual" dur="indefinite" 
	region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" src="medias/example.svg"/>
</seq>
```


###### Attribute *@begin*
This attribute defines the beginning of the sequence.  
"0s" means the sequence starts to play as the same time as the PAR element starts.  
"3s" means the sequence starts to play 3 seconds after the beginning of PAR element.

###### Attribute *@repeatCount*
This attribute defines the number of times to repeat the sequence.  
"indefinite" means the sequence is played in loop.  
 
###### Attribute *@region* 
This attribute defines the identifier of the region where the sequence must be played.  
The position and the style of the region are described in the STYLE element.

###### Child elements
The child elements are the medias to play.  
Each media is played one after one, like a sequence. 
 
```
 <img class="visual" dur="00:00:03.000" 
    region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" src="medias/yellow.jpg"/>
```
```
<img class="visual" dur="indefinite" 
    region="UIDe9eb19c6_f089_4d5e_b1db_59e5a201ca63" src="medias/example.svg"/>
```
The attribute *@dur* defines the duration for playing the media.  
"00:00:03.000" means the media will be played during 3 seconds.  
"indefinite" means there is no time limit.






