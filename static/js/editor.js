
(function() {
  var editorNode = document.getElementById('editor');
  var editorParentNode = editorNode.parentNode;
  
  var editor = ace.edit("editor");
  editor.setShowPrintMargin(false);
  editor.setTheme("ace/theme/tomorrow");
  editor.getSession().setUseWorker(false);
  editorNode.style.width = editorParentNode.clientWidth + "px"; 
  
  var JavaScriptMode = ace.require("ace/mode/javascript").Mode;
  editor.session.setMode(new JavaScriptMode());
})();
