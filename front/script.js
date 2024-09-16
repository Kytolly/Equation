function createInputFields() {
    var n = parseInt(document.getElementById("unknowns").value);
    if (n < 1) {
        n = 1;
    } else if (n > 6) {
        n = 6;
    }

    var inputFields = "";
    for (var i = 1; i <= n; i++) {
        for (var j = 1; j <= n; j++) {
            inputFields += "<input type='text' style='width:40px' maxlength='10' class='matrix-element' id='input-x" + i + j + "'>" + "x" + j;
            // inputFields += "x" + i + j + ": <input type='text' class='matrix-element' id='input-x" + i + j + "'> ";
            if (j != n) {
                inputFields += "+";
            }
        }
        var j = n + 1;
        inputFields += "=" + "<input type='text style='width:40px' maxlength='10'class='matrix-element' id = 'input-x" + i + j + "'>";
        inputFields += "<br>";
    }
    document.getElementById("inputFields").innerHTML = inputFields;
}

function solveEquation() {
    var n = parseInt(document.getElementById("unknowns").value);
    if (n < 1) { n = 1; }
    if (n > 6) { n = 6; }

    var matrix = [];
    var expansion = [];
    for (var i = 1; i <= n; i++) {
        var row = [];
        for (var j = 1; j <= n; j++) {
            var value = parseFloat(document.getElementById("input-x" + i + j).value);
            row.push(value);
        }
        var j = n + 1;
        var value = parseFloat(document.getElementById("input-x" + i + j).value);
        expansion.push(value);
        matrix.push(row);
    }


    console.log(JSON.stringify({ n: n, matrix: matrix, expansion: expansion }));


    $.ajax({
        type: "POST",
        url: "http://localhost:8080/front",
        data: JSON.stringify({ n: n, matrix: matrix, expansion: expansion }),
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            var result = response;  // 提取返回的结果
            document.getElementById("result").innerHTML = result;
        },
        error: function (error) {
            console.error("ajax request fail:", textStatus, errorThrown);
            document.getElementById("result").innerHTML = "请求失败了～";
        }
    });
}