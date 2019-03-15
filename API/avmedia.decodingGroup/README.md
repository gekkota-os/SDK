#  Video decoding group

The video decoding group is a CSS property allowing to share decoding instances between several HTML video elements.
Every element with the same property value is sharing **one** decoding instance.

**Requirement**: the preference "innes.video.decoding-group.enabled" must be enabled or the CSS property will not be considered.

| Property name  | type | value |
|----------------|------|-------|
| -gkt-video-decoding-group | String   | Any string value |

Use it like any other CSS property:

````css
#html_element_1 {
    position: absolute;
    left: 0px;
    width: 100%;
    -gkt-video-decoding-group:"decoding_group_name";
}

#html_element_2 {
    position: relative;
    left: 20px;
    height: 100%;
    -gkt-video-decoding-group:"decoding_group_name";
}
````

In this example, if "decoding_group_name" is the same in "html_element_1" and "html_element_2", both elements are sharing one decoding instance. An HTML element with no -gkt-video-decoding-group specified or with only itself in the decoding group will have the same behavior as a default HTML video element.
