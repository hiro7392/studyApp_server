<head>
  <meta charset="UTF-8">
  <title>講師用画面</title>
  <link href="../client/css/teacher.css" rel="stylesheet">
</head>

<body>
  <header>
    <a class="header-logo" href="teacher.html">アプリ名</a>
  </header>
  <div class="teacher-wrapper">
    <!-- 検索条件に合う生徒を探して、次のページ(search_student.html)で一覧を表示する -->
    <div class="container">
      <div class="block">
        <h1>生徒検索</h1>
        <form class="form-element-wrapper" action="/search_student">
          <input type="text" name="student_name" class="form-control" placeholder="キーワード">
          <input type=submit value="検索">
        </form>
        <h3>絞り込み</h3>
        <form action="/show_all_student/">
          <div class="searcher-detail-item">
            <span class="searcher-detail-item-header">学年　　　 : </span>
            <input type="checkbox" id="grade" name="grade">
            <label for="scales">１年</label>
            <input type="checkbox" id="grade" name="grade">
            <label for="scales">２年</label>
            <input type="checkbox" id="grade" name="grade">
            <label for="scales">３年</label>
          </div>
          </form>
        <div class="searcher-detail-item">
        
          
      <form action="/show_all_student/">
      <span class="searcher-detail-item-header" >担当講師名 : </span>
      <input type="text" name="teacher_name" class="form-control" placeholder="担当講師名" autofocus>
      <input type="submit" value="表示">
    </form>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="block">
        <!-- 入力したデータを次のページで表示して確認する -->
        <h1>生徒登録</h1>
        <form action="/register_student/">
          <ul>
            <li> 学年　　 :
              <input type="text" name="grade" class="form-control">
            </li>
            <li> 生徒名　 :
              <input type="text" name="student_name" class="form-control">
            </li>
            <li> 担当講師 :
              <input type="text" name="teacher_name" class="form-control">
            </li>
            <li> 中学校　 :
              <input type="text" name="school" class="form-control">
            </li>
            <li> 志望校　 :
              <input type="text" name="school_of_choise" class="form-control">
            </li>
          </ul>
          <input type=submit class="right" value="登録">
        </form>
      </div>
      <div class="block">
        <!-- 入力したデータを次のページで表示して確認する -->
      

        <h1>教材登録</h1>
        <form action="/register_textbook/">
          <ul>
            <li> タイトル :
              <input type="text" name="text_title" class="form-control">
            </li>
            <li> 科目　　 :
              <input type="text" name="subject" class="form-control">
            </li>
            <li> ページ数 :
              <input type="text" name="pages" class="form-control">
            </li>
          </ul>
          <input type=submit class="right" value="登録">
        </form>
      </div>
    </div>
  </div>
</body>
