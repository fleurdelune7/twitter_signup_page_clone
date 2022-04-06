{{template "base" .}} 

{{define "title"}} Sign Up {{end}}


{{define "style"}}     <link rel="stylesheet" type="text/css" href="/static/css/style.css" />{{end}}
{{define "content"}}
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="style.css">
    <title>Sign Up for Twitter</title>
</head>
<body>
<div>
<div class="center">
<div class = "up1">
    <img class = "x" src="static/css/images/x.png" alt="x">
    <img class = "logo" src="static/css/images/logo.png" alt="Logo">
</div>
<div class = "up2">
    <h2 class="cya">Create your account</h2>
</div>
<form action="/users" method="post">
<div class="up3">
    <input type="text" class="name" id="first_name" name="first_name" autocomplete="off" placeholder="Name">
    <input type="text" class="name" id="email" name="email" autocomplete="off" placeholder="Email">
</div>
<div class="up4">
    <p class="upi">Use phone instead</p>
</div>
<div class="up5">
    <p class="dob">Date of birth</p>
    <p class="dobp">This will not be shown publicly. Confirm your own age, even if this account is for a business, a pet, or something else.</p>
</div>
<div class="up6">
    <select id="month" name="month">
        <option>Month</option>
        <option value="january">January</option>
        <option value="february">February</option>
        <option value="march">March</option>
        <option value="april">April</option>
        <option value="may">May</option>
        <option value="june">June</option>
        <option value="july">July</option>
        <option value="august">August</option>
        <option value="september">September</option>
        <option value="october">October</option>
        <option value="november">November</option>
        <option value="december">December</option>
    </select>
    <select id="day" name="day">
        <option>Day</option>
        <option value="01">01</option>
        <option value="02">02</option>
        <option value="03">03</option>
        <option value="04">04</option>
        <option value="05">05</option>
        <option value="06">06</option>
        <option value="07">07</option>
        <option value="08">08</option>
        <option value="09">09</option>
        <option value="10">10</option>
        <option value="11">11</option>
        <option value="12">12</option>
        <option value="13">13</option>
        <option value="14">14</option>
        <option value="15">15</option>
        <option value="16">16</option>
        <option value="17">17</option>
        <option value="18">18</option>
        <option value="19">19</option>
        <option value="20">20</option>
        <option value="21">21</option>
        <option value="22">22</option>
        <option value="23">23</option>
        <option value="24">24</option>
        <option value="25">25</option>
        <option value="26">26</option>
        <option value="27">27</option>
        <option value="28">28</option>
        <option value="29">29</option>
        <option value="30">30</option>
        <option value="31">31</option>
    </select>
    <select id="year" name="year">
        <option>Year</option>
        <option value="1960">1960</option>
        <option value="1961">1961</option>
        <option value="1962">1962</option>
        <option value="1963">1963</option>
        <option value="1964">1964</option>
        <option value="1965">1965</option>
        <option value="1966">1966</option>
        <option value="1967">1967</option>
        <option value="1968">1968</option>
        <option value="1969">1969</option>
        <option value="1970">1970</option>
        <option value="1971">1971</option>
        <option value="1972">1972</option>
        <option value="1973">1973</option>
        <option value="1974">1974</option>
        <option value="1975">1975</option>
        <option value="1976">1976</option>
        <option value="1977">1977</option>
        <option value="1978">1978</option>
        <option value="1979">1979</option>
        <option value="1980">1980</option>
        <option value="1981">1981</option>
        <option value="1982">1982</option>
        <option value="1983">1983</option>
        <option value="1984">1984</option>
        <option value="1985">1985</option>
        <option value="1986">1986</option>
        <option value="1987">1987</option>
        <option value="1988">1988</option>
        <option value="1989">1989</option>
        <option value="1990">1990</option>
        <option value="1991">1991</option>
        <option value="1992">1992</option>
        <option value="1993">1993</option>
        <option value="1994">1994</option>
        <option value="1995">1995</option>
        <option value="1996">1996</option>
        <option value="1997">1997</option>
        <option value="1998">1998</option>
        <option value="1999">1999</option>
        <option value="2000">2000</option>
        <option value="2001">2001</option>
        <option value="2002">2002</option>
        <option value="2003">2003</option>
        <option value="2004">2004</option>
        <option value="2005">2005</option>
        <option value="2006">2006</option>
        <option value="2007">2007</option>
        <option value="2008">2008</option>
        <option value="2009">2009</option>
        <option value="2010">2010</option>
        <option value="2011">2011</option>
        <option value="2012">2012</option>
        <option value="2013">2013</option>
        <option value="2014">2014</option>
        <option value="2015">2015</option>
        <option value="2016">2016</option>
        <option value="2017">2017</option>
        <option value="2018">2018</option>
        <option value="2019">2019</option>
        <option value="2020">2020</option>
        <option value="2021">2021</option>
        <option value="2022">2022</option>
    </select>
</div>
<div class="up7">
    <input type="submit" value="Sign Up" class="Next">
</div>
</form>
</div>
</div>

{{if .Form}}
<ul class="error">
{{ range $key, $value := .Form.Errors }}
   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
{{ end }}
<ul>
{{end}}

{{end}}



