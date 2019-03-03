function varBehavior() {
    var v = "AAA";
    console.log(v)
    {
        var v = "BBB";
        console.log(v)
    }
    console.log(v)
    newLine()
}

function letBehavior() {
    let v = "aaa";
    console.log(v)
    {
        let v = "bbb";
        console.log(v)
    }
    console.log(v)
    newLine()
}

/*
function constBehavior() {
    const v = "ccc";
    console.log(v);
    v = "CCC";
    console.log(v);

    newLine();
}
*/

varBehavior();
letBehavior();
//constBehavior();


function newLine() {
    console.log()
}