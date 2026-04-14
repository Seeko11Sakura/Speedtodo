import '../../tasks/presentation/task_list_page.dart';
import 'package:flutter/widgets.dart';

class UpcomingPage extends StatelessWidget {
  const UpcomingPage({super.key});

  @override
  Widget build(BuildContext context) => const TaskListPage(view: 'upcoming');
}
