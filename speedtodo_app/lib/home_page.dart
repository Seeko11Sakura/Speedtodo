import 'package:flutter/material.dart';
import 'today_page.dart';
import 'inbox_page.dart';
import 'calendar_page.dart';
import 'pomodoro_page.dart';
import 'lists_page.dart';
import 'tags_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _index = 0;

  final List<_NavTab> _tabs = const [
    _NavTab(icon: Icons.today, label: 'Today'),
    _NavTab(icon: Icons.inbox, label: 'Inbox'),
    _NavTab(icon: Icons.calendar_month, label: 'Calendar'),
    _NavTab(icon: Icons.timer, label: 'Focus'),
    _NavTab(icon: Icons.list_alt, label: 'Lists'),
    _NavTab(icon: Icons.label, label: 'Tags'),
  ];

  final List<Widget> _pages = const [
    TodayPage(),
    InboxPage(),
    CalendarPage(),
    PomodoroPage(),
    ListsPage(),
    TagsPage(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_index],
      bottomNavigationBar: NavigationBar(
        selectedIndex: _index,
        onDestinationSelected: (i) => setState(() => _index = i),
        destinations: _tabs
            .map((t) => NavigationDestination(icon: Icon(t.icon), label: t.label))
            .toList(),
      ),
    );
  }
}

class _NavTab {
  const _NavTab({required this.icon, required this.label});
  final IconData icon;
  final String label;
}
