package lib

import (
	"sort"
)

// Grouper is the group struct.
type Grouper struct{}

// GroupedTodos is the main struct of this file.
type GroupedTodos struct {
	Groups map[string][]*Todo
}

// GroupByTag is grouping todos by its tag.
func (g *Grouper) GroupByTag(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}

	allTags := []string{}

	for _, todo := range todos {
		allTags = AddIfNotThere(allTags, todo.Tags)
	}

	for _, todo := range todos {
		for _, tag := range todo.Tags {
			groups[tag] = append(groups[tag], todo)
		}
		if len(todo.Tags) == 0 {
			groups["No tags"] = append(groups["No tags"], todo)
		}
	}

	// finally, sort the todos
	for groupName, todos := range groups {
		groups[groupName] = g.sort(todos)
	}

	return &GroupedTodos{Groups: groups}
}

// GroupByProject is grouping todos by its project.
func (g *Grouper) GroupByProject(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}

	allProjects := []string{}

	for _, todo := range todos {
		allProjects = AddIfNotThere(allProjects, todo.Projects)
	}

	for _, todo := range todos {
		for _, project := range todo.Projects {
			groups[project] = append(groups[project], todo)
		}
		if len(todo.Projects) == 0 {
			groups["No projects"] = append(groups["No projects"], todo)
		}
	}

	// finally, sort the todos
	for groupName, todos := range groups {
		groups[groupName] = g.sort(todos)
	}

	return &GroupedTodos{Groups: groups}
}

// GroupByStatus is grouping todos by status
func (g *Grouper) GroupByStatus(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}

	for _, todo := range todos {
		if todo.Status != "" {
			groups[todo.Status] = append(groups[todo.Status], todo)
		} else {
			groups["No status"] = append(groups["No status"], todo)
		}
	}

	// finally, sort the todos
	for groupName, todos := range groups {
		groups[groupName] = g.sort(todos)
	}

	return &GroupedTodos{Groups: groups}
}

// GroupByNothing is the default result if todos are not grouped by tag project.
func (g *Grouper) GroupByNothing(todos []*Todo) *GroupedTodos {
	groups := map[string][]*Todo{}
	groups["all"] = g.sort(todos)

	return &GroupedTodos{Groups: groups}
}

func (g *Grouper) sort(todos []*Todo) []*Todo {
	sort.Slice(todos, func(i, j int) bool {
		// always favor unarchived todos over archived ones
		if todos[i].Archived || todos[j].Archived {
			return !todos[i].Archived || todos[j].Archived
		}

		// always favor un-completed todos over completed ones
		if todos[i].Completed || todos[j].Completed {
			return !todos[i].Completed || todos[j].Completed
		}

		// always favor prioritized todos
		if todos[i].IsPriority || todos[j].IsPriority {
			return todos[i].IsPriority || !todos[j].IsPriority
		}

		// always prefer a todo with a due date
		if todos[i].Due != "" && todos[j].Due == "" {
			return true
		}

		// un-favor todos without a due date
		if todos[i].Due == "" {
			return false
		}

		return todos[i].CalculateDueTime().Before(todos[j].CalculateDueTime())
	})

	return todos
}
