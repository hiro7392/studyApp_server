<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒情報)</title>
  <link href="../client/css/teacher_student.css" rel="stylesheet">
</head>

<body>

<header>
<a class="a_button" href="teacher.html">ホーム</a>
</header>

<h2>生徒情報</h2>
<ul>
  <li>氏名 : {{.St.Name}}</li>
  <li>学年 : {{.St.Grade}}</li>
  <li>中学校名 :{{.St.NowSchool}}</li>
  <li>志望校 : {{.St.WantSchool}}</li>
</ul>

<h2>お知らせ</h2>
<div>
  今週は
  {{range $nowdata:=.Tas}}
  <ul>
    <li>{{$nowdata.TextName}} の {{$nowdata.Nowpage}} ~ {{$nowdata.Allpage}}</li>
  </ul>
  と
  {{end}}
  を進めてください。
</div>

<h2>ガントチャート</h2>
<div class="wrapper">
  <div class="gantt">
    <div class="gantt__row gantt__row--months">
      <div class="gantt__row-first">
          月
      </div>
      <span> 4月</span><span></span><span></span><span></span>
      <span> 5月</span><span></span><span></span><span></span>
      <span> 6月</span><span></span><span></span><span></span>
      <span> 7月</span><span></span><span></span><span></span>
      <span> 8月</span><span></span><span></span><span></span>
      <span> 9月</span><span></span><span></span><span></span>
      <span>10月</span><span></span><span></span><span></span>
      <span>11月</span><span></span><span></span><span></span>
      <span>12月</span><span></span><span></span><span></span>
      <span> 1月</span><span></span><span></span><span></span>
      <span> 2月</span><span></span><span></span><span></span>
      <span> 3月</span><span></span><span></span><span></span>
    </div>
    <div class="gantt__row gantt__row--lines" data-month="5">
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span class="marker"></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
      <span></span><span></span><span></span><span></span>
    </div>
    <div class="gantt__row">
      <div class="gantt__row-first">
        国語
      </div>
      <ul class="gantt__row-bars">
        {{range $taskData:=.Taskbox.TaskJapa}}
        <li style="grid-column: {{$taskData.Start}}/{{$taskData.End}}; background-color: green;">{{$taskData.Name}}</li>
        {{end}}
      </ul>
    </div>
    <div class="gantt__row gantt__row">
      <div class="gantt__row-first">
        数学
      </div>
      <ul class="gantt__row-bars">
      {{range $taskData:=.Taskbox.TaskMath}}
        <li style="grid-column: {{$taskData.Start}}/{{$taskData.End}}; background-color: black;">{{$taskData.Name}}</li>
      {{end}}
      </ul>
    </div>
    <div class="gantt__row">
      <div class="gantt__row-first">
        理科
      </div>
      <ul class="gantt__row-bars">
        {{range $taskData:=.Taskbox.TaskSci}}
        <li style="grid-column: {{$taskData.Start}}/{{$taskData.End}}; background-color: red;">{{$taskData.Name}}</li>
        {{end}}
      </ul>
    </div>
    <div class="gantt__row gantt__row">
      <div class="gantt__row-first">
        社会
      </div>
      <ul class="gantt__row-bars">
        {{range $taskData:=.Taskbox.TaskSoc}}
        <li style="grid-column: {{$taskData.Start}}/{{$taskData.End}}; background-color: blue;">{{$taskData.Name}}</li>
        {{end}}
      </ul>
    </div>
    <div class="gantt__row gantt__row">
      <div class="gantt__row-first">
        英語
      </div>
      <ul class="gantt__row-bars">
        {{range $taskData:=.Taskbox.TaskEng}}
        <li style="grid-column: {{$taskData.Start}}/{{$taskData.End}}; background-color: orange;">{{$taskData.Name}}</li>
        {{end}}
      </ul>
    </div>
  </div>

<footer>
  <a class="a_button" href="info_student.html">生徒資料作成</a>
</footer>
</body>
