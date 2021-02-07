package store

// Submission is a post DB type
type Submission struct {
	URL string
	ID  string
	UID string
}

// Submissions is a list of all submissions
var Submissions []*Submission

// AddSubmission adds a post to submissions list
func AddSubmission(post *Submission) {
	Submissions = append(Submissions, post)
}

// RemoveSubmissions removes all submissions
func RemoveSubmissions() {
	Submissions = nil
}
