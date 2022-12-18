// ref: http://www.webclap-dandy.com/?category=Programing&id=5
const tabKey = 9;
const spaces = '  ';
const onTabKey = (e) => {
  if (e.keyCode != tabKey){ return; };

  e.preventDefault();

  const cursor = e.target;
  const currStr = String(cursor.value);

  const position = cursor.selectionStart;
  const left = currStr.substring(0, position);
  const right = currStr.substring(position, currStr.length);

  cursor.value = left + spaces + right;
  cursor.selectionEnd = position + spaces.length;
};
document.getElementById('in').onkeydown = (e) => onTabKey(e);
