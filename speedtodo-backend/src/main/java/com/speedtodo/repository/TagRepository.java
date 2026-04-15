package com.speedtodo.repository;
import com.speedtodo.model.Tag;
import org.springframework.data.jpa.repository.JpaRepository;
public interface TagRepository extends JpaRepository<Tag, Long> {}