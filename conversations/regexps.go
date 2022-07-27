package conversations

import "regexp"

var regSendMyUsername, _ = regexp.Compile("[0-9]+..sendMyUsername")

var regLike, _ = regexp.Compile("[0-9]+..like")

var regRepeatSendingUsername, _ = regexp.Compile("[0-9]+..repeatSendingUsername")

var regFeedback, _ = regexp.Compile("[0-9]+..answerFeedback")

var regNumVerifier = regexp.MustCompile("[0-9]+")
