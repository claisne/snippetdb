
(function() {
  var editorNode = document.getElementById('editor');
  var editorParentNode = editorNode.parentNode;
  editorNode.style.width = editorParentNode.clientWidth + "px"; 
  var editor = ace.edit("editor");
  var JavaScriptMode = ace.require("ace/mode/javascript").Mode;
  console.log(ace.require("ace/mode/javascript"));
  editor.session.setMode(new JavaScriptMode());
})();
