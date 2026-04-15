package com.speedtodo.repository;
import com.speedtodo.model.TaskList;
import org.springframework.data.jpa.repository.JpaRepository;
public interface TaskListRepository extends JpaRepository<TaskList, Long> {}