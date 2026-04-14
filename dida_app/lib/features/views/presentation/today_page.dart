import '../../tasks/presentation/task_list_page.dart';
import 'package:flutter/widgets.dart';

class TodayPage extends StatelessWidget {
  const TodayPage({super.key});

  @override
  Widget build(BuildContext context) => const TaskListPage(view: 'today');
}
