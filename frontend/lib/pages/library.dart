import 'package:flutter/material.dart';
import 'package:frontend/entities/languageInfo.dart';
import 'package:frontend/services/languageInfo.dart';
import 'package:frontend/pages/components/languageInfo.dart';
import 'package:frontend/pages/components/sideBar.dart';

class LibraryScreen extends StatefulWidget {
  LibraryScreen({Key key}) : super(key: key);

  @override
  _LibraryScreenState createState() => _LibraryScreenState();
}

class _LibraryScreenState extends State<LibraryScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Library"),
      ),
      drawer: SideBar(),
      body: Center(
        child: FutureBuilder<List<LanguageInfo>>(
          future: fetchLanguageInfo(),
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return LanguageInfoList(languageList: snapshot.data);
            } else if (snapshot.hasError) {
              return Text("${snapshot.error}");
            }
            // By default, show a loading spinner.
            return CircularProgressIndicator();
          },
        ),
      ),
    );
  }
}
