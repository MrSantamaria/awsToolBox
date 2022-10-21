package awsToolBox

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func (AwsSession *awsSession) Iam() error {
	if AwsSession.session == nil {
		err := AwsSession.InitialLogin()
		if err != nil {
			return err
		}
	}

	//AwsSession.ListOpenIDConnectProviders()
	//AWSSession.ListRoles()
	AwsSession.ListPolicies()

	return nil
}

func (AwsSession *awsSession) ListOpenIDConnectProviders() error {
	input := &iam.ListOpenIDConnectProvidersInput{}
	result, err := AwsSession.iam.ListOpenIDConnectProviders(input)
	if err != nil {
		return err
	}

	for _, provider := range result.OpenIDConnectProviderList {
		fmt.Println(*provider.Arn)
		//Create GetOpenIDConnectProviderInput struct
		input := &iam.GetOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: provider.Arn,
		}

		//Get the provider
		result, err := AwsSession.iam.GetOpenIDConnectProvider(input)
		if err != nil {
			return err
		}

		//Add flags to use a name for the provider and days to expire
		//If the name contains "rh-oidc" and is older than 90 days, delete it
		if strings.Contains(*result.Url, "rh-oidc") && time.Since(*result.CreateDate) > 90*24*time.Hour {
			//Create DeleteOpenIDConnectProviderInput struct
			input := &iam.DeleteOpenIDConnectProviderInput{
				OpenIDConnectProviderArn: provider.Arn,
			}

			//Delete the provider
			_, err := AwsSession.iam.DeleteOpenIDConnectProvider(input)
			if err != nil {
				return err
			}

			fmt.Println("Deleted provider: ", *provider.Arn)
		}
	}

	return nil
}

func (AwsSession *awsSession) ListRoles() error {
	input := &iam.ListRolesInput{
		MaxItems: aws.Int64(1000),
	}
	result, err := AwsSession.iam.ListRoles(input)
	if err != nil {
		return err
	}

	for _, role := range result.Roles {
		if strings.Contains(*role.Arn, "osde2e") && time.Since(*role.CreateDate) > 90*24*time.Hour {
			fmt.Printf("Attempting to delete role: %s\n", *role.RoleName)
			//Remove Roles from Instance Profiles
			//Create ListInstanceProfilesForRoleInput struct
			instanceProfilesForRoleInputinput := &iam.ListInstanceProfilesForRoleInput{
				RoleName: role.RoleName,
			}

			//Get the instance profiles
			instanceProfiles, errInstanceProfiles := AwsSession.iam.ListInstanceProfilesForRole(instanceProfilesForRoleInputinput)
			if errInstanceProfiles != nil {
				return fmt.Errorf("error getting instance profiles for role: %s", errInstanceProfiles)
			}

			//Remove the roles from the instance profiles
			for _, instanceProfile := range instanceProfiles.InstanceProfiles {
				//Create RemoveRoleFromInstanceProfileInput struct
				removeRoleFromInstanceProfileInput := &iam.RemoveRoleFromInstanceProfileInput{
					InstanceProfileName: instanceProfile.InstanceProfileName,
					RoleName:            role.RoleName,
				}

				//Remove the role from the instance profile
				_, errRemoveRoleFromInstanceProfile := AwsSession.iam.RemoveRoleFromInstanceProfile(removeRoleFromInstanceProfileInput)
				if errRemoveRoleFromInstanceProfile != nil {
					return fmt.Errorf("error removing role from instance profile: %s", errRemoveRoleFromInstanceProfile)
				}

				//Need to create a wait function to wait for the role to be removed from the instance profile
				fmt.Println("Removed role from instance profile: ", *instanceProfile.Arn)
			}

			//Delete policy inline to the role
			inlineRolePoliciesInput := &iam.ListRolePoliciesInput{
				RoleName: role.RoleName,
			}
			inlinePolicies, errInlineRolePoliciesInput := AwsSession.iam.ListRolePolicies(inlineRolePoliciesInput)
			if errInlineRolePoliciesInput != nil {
				return errInlineRolePoliciesInput
			}

			for _, policy := range inlinePolicies.PolicyNames {
				input := &iam.DeleteRolePolicyInput{
					PolicyName: policy,
					RoleName:   role.RoleName,
				}

				_, errInlinePolicies := AwsSession.iam.DeleteRolePolicy(input)
				if errInlinePolicies != nil {
					return fmt.Errorf("error deleting inline policy: %s", errInlinePolicies)
				}

				fmt.Println("Deleted inline policy: ", *policy)
			}

			//Delete policy attached to the role
			attachedRolePoliciesInput := &iam.ListAttachedRolePoliciesInput{
				RoleName: role.RoleName,
			}
			attachedPolicies, errAttachedRolePoliciesInput := AwsSession.iam.ListAttachedRolePolicies(attachedRolePoliciesInput)
			if errAttachedRolePoliciesInput != nil {
				return errAttachedRolePoliciesInput
			}

			for _, policy := range attachedPolicies.AttachedPolicies {
				detachInput := &iam.DetachRolePolicyInput{
					PolicyArn: policy.PolicyArn,
					RoleName:  role.RoleName,
				}

				_, errAttachedPolicies := AwsSession.iam.DetachRolePolicy(detachInput)
				if errAttachedPolicies != nil {
					return errAttachedPolicies
				}

				// //Delete the policy
				// deleteInput := &iam.DeletePolicyInput{
				// 	PolicyArn: policy.PolicyArn,
				// }

				// _, errDeletePolicy := AwsSession.iam.DeletePolicy(deleteInput)
				// if errDeletePolicy != nil {
				// 	return fmt.Errorf("error deleting policy: %s", errDeletePolicy)
				// }

				// time.Sleep(10 * time.Second)
				// fmt.Println("Deleted attached policy: ", *policy.PolicyArn)
			}

			roleInput := &iam.DeleteRoleInput{
				RoleName: role.RoleName,
			}

			//Delete the role
			_, err = AwsSession.iam.DeleteRole(roleInput)
			if err != nil {
				return err
			}

			fmt.Println("Deleted role: ", *role.Arn)
		}
	}

	return nil
}

func (AwsSession *awsSession) ListPolicies() error {
	input := &iam.ListPoliciesInput{
		MaxItems: aws.Int64(1000),
	}
	result, err := AwsSession.iam.ListPolicies(input)
	if err != nil {
		return err
	}

	for _, policy := range result.Policies {
		if strings.Contains(*policy.Arn, "osde2e") && time.Since(*policy.CreateDate) > 90*24*time.Hour {
			fmt.Printf("Attempting to delete policy: %s", *policy.Arn)
			input := &iam.DeletePolicyInput{
				PolicyArn: policy.Arn,
			}

			//Delete the policy
			_, err := AwsSession.iam.DeletePolicy(input)
			if err != nil {
				return err
			}

			fmt.Println("Deleted policy: ", *policy.Arn)
		}
	}

	return nil
}
