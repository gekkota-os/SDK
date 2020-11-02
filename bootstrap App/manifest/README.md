# GEKKOTA Bootstrap App Manifest

## Introduction to Manifest file

This document describes the content of the Manifest file of Gekkota bootstrap App.
The format of the Manifest file is RDF XML.
When the file is retrieved in the "Pull" mode by Gekkota using a webdav server, its naming is as following :
*manifest.[ID].xml*, where *[ID]* is the unique identifier of the device .

This file allows to :

* Embbed some metadata coming from a CMS (metadata),
* Define the point of entry to launch the App of startup (launcher),
* Enumerate all the files required to run the application (cache).

````mermaid
graph TD
    B[manifest file]
    B-->C(metadata)
    B-->D(launcher);
    B-->E(cache);
````

## Example

````xml
<RDF xmlns="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:pzpm="ns.innes.gekkota.manifest" xmlns:cms="ns.innes.example">
  <Description about=" ns.innes.gekkota.manifest#metadata">
    <!--all metadatas inserted here will be putted in devices status file-->
    <pzpm:publish-id>e5fdaa88-7eb3-4e20-8602-1d94030d2bce</pzpm:publish-id>
    <pzpm:publish-size>654681689</pzpm:publish-size>
    <pzpm:publish-generator>Screen composer 3.10.16</pzpm:publish-generator>
    <pzpm:publish-date>2013-10-16T14:41:50.783Z</pzpm:publish-date>
    <cms:playout-id>e5fdaa88-7eb3-4e20-8602-1d94030d2bc6</cms:playout-id>
    <!--others metadatas...-->
  </Description>

  <!--Configuration of launcher-->
  <Description about=" ns.innes.gekkota.manifest#launcher">
    <pzpm:bootstrap src="player.html"/>
  </Description>

  <!--List of resources-->
  <Description about=" ns.innes.gekkota.manifest#cache">
    <Bag>
      <li>.medias/videos/1280-720-25-p-high-3.1-90_counter.mp4</li>
      <li src="http://monServeur/fichier.jpg" username="username" password="password" refresh="120">.medias/remote.jpg</li>
      <li src="urn:innes:storage:removable/videos/">.domain-repository/videos/</li>
      <li src="urn:innes:storage:local">.domain-repository/audios/</li>
      <li src="urn:innes:storage:removable:/subtitles/fichier.srt">.domain-repository/subtitles/fichier.srt</li>
      <li>player.html</li>
      <li>variables.xml</li>
    </Bag>
  </Description>
</RDF>
````

## Section <*#metadata*>

This section is defined by a <*Description*> RDF element which attribute *@about* is always set to "*ns.innes.gekkota.manifest#metadata*".
For retrocompatibility, *"ns.innes.playzilla.manifest#metadata"* and *"urn:innes:manifest:metadata"* are authorized.
This section contains the metadata of the App of startup. This element is optional. If present, it contains few required metadata. The other metadata are free.

### Element <*publish-id*>

Contain an id for publication which is an uuid.
This element is required if <*#metadata*> is present.

### Element <*publish-size*>

Contain the size in bytes of the publication (excluding the manifest.xml file and external resources). This element is required if <*#metadata*> is present.

### Element <*publish-generator*>

Contain the name of the generator of the Manifest file, which is a choosen string. This element is required if <*#metadata*> is present.

### Element <*publish-date*>

Contain the date of publication, following ISO-8601 format. This element is required if <*#metadata*> is present.

## Section <*#launcher*>

This section is defined by a <*Description*> RDF element which attribute *@about* is always set to *"ns.innes.gekkota.manifest#launcher"*
For retrocompatibility, *"ns.innes.playzilla.manifest#launcher"* and *"urn:innes:manifest:launcher"* are authorized.
This section allows to configure the Gekkota "launcher" component.
It is optional.

### Element <*bootstrap*>

This element is required when section <*#launcher*> is present.

#### Attribute *@src*

*@src* contains the URI of the html document to load. If the URI is relative, then it is relative to the Manifest file.
If the <*#launcher*> section is present, then this element is required, else the default value of the *@src* attribute is *player.html*.
The *@username* and *@password* attributes are optional and defined for authentication.

## Section *<#cache>*

This section is defined by a <*Description*> RDF element which attribute "@about" is always set to "ns.innes.gekkota.manifest#cache".
For retrocompatibility, *"ns.innes.playzilla.manifest#cache"* and *"urn:innes:manifest:cache"* are authorized.
This element is a RDF Bag and contains the list of all the resources required to run  the App of startup.
This element is required.
Every URIs in the Bag must be relative to the Manifest file.

### Element <*Bag*>

This element contains a list of RDF <*li*> elements.

#### Attribute *@refresh*

The presence of the *@refresh* attribute implies that the resource will be repeatedly downloaded through webdav protocol with a time interval matching the value bound to the attribute in seconds.

#### Attribute *@keep-only*

The *@keep-only* attribute indicates that the file bound to the <*li*> element must be taken in count only by the purger : the file is not downloaded in Gekkota. The presence of such a file is not required to validate the publication of the Manifest file.

#### Attribute *@username*

The *@username* attribute specifies the login for the webdav connection.

#### Attribut *@password*

The *@password* attribute specifies the password for the webdav connection.
